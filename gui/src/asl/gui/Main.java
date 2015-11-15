/*
 * The MIT License
 *
 * Copyright 2015 Ozan Egitmen.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */
package asl.gui;

import java.awt.Color;
import java.io.BufferedReader;
import java.io.File;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.prefs.Preferences;
import javax.swing.BorderFactory;
import javax.swing.JFileChooser;
import javax.swing.JTextField;
import javax.swing.SwingUtilities;
import javax.swing.UIManager;
import javax.swing.UnsupportedLookAndFeelException;
import javax.swing.border.Border;
import javax.swing.filechooser.FileNameExtensionFilter;
import javax.swing.plaf.ColorUIResource;

public class Main extends javax.swing.JFrame {

    Preferences prefs = Preferences.userRoot().node(this.getClass().getName());
    boolean aslFix, outputDirFix, gaveError = false;

    public Main() {
        initComponents();
        getContentPane().setBackground(Color.WHITE);
        lblASLError.setText(" ");
        lblInputError.setText(" ");
        lblOutputError.setText(" ");
        txtASLDir.setText(prefs.get("aslDir", ""));
        txtInputDir.setText(prefs.get("inputDir", ""));
        txtOutputDir.setText(prefs.get("outputDir", ""));
        cbCompileAll.setSelected(prefs.getBoolean("compileAll", false));
        cbPrettyPrinting.setSelected(prefs.getBoolean("prettyPrinting", false));
    }

    private String fileChooser(String title, FileNameExtensionFilter fileType) {
        JFileChooser chooser = new JFileChooser();
        if (fileType != null) {
            chooser.setFileFilter(fileType);
            chooser.setAcceptAllFileFilterUsed(false);
            chooser.setFileSelectionMode(0);
        } else {
            chooser.setFileSelectionMode(1);
        }
        chooser.setDialogTitle(title);
        String selectedPath = "";
        if (chooser.showOpenDialog(null) == 0) {
            selectedPath = chooser.getSelectedFile().toString();
        } else {
            chooser.cancelSelection();
        }
        return selectedPath;
    }

    @SuppressWarnings("unchecked")
    // <editor-fold defaultstate="collapsed" desc="Generated Code">//GEN-BEGIN:initComponents
    private void initComponents() {

        lblInput = new javax.swing.JLabel();
        txtInputDir = new javax.swing.JTextField();
        lblOutput = new javax.swing.JLabel();
        txtOutputDir = new javax.swing.JTextField();
        btnInput = new javax.swing.JButton();
        btnOutput = new javax.swing.JButton();
        lblASL = new javax.swing.JLabel();
        txtASLDir = new javax.swing.JTextField();
        btnASL = new javax.swing.JButton();
        jSeparator = new javax.swing.JSeparator();
        lblASLSmall = new javax.swing.JLabel();
        lblInputSmall = new javax.swing.JLabel();
        lblOutputSmall = new javax.swing.JLabel();
        cbCompileAll = new javax.swing.JCheckBox();
        cbPrettyPrinting = new javax.swing.JCheckBox();
        btnCompile = new javax.swing.JButton();
        lblASLError = new javax.swing.JLabel();
        lblInputError = new javax.swing.JLabel();
        lblOutputError = new javax.swing.JLabel();

        setDefaultCloseOperation(javax.swing.WindowConstants.EXIT_ON_CLOSE);
        setTitle("ASL GUI");
        setResizable(false);

        lblInput.setFont(new java.awt.Font("Microsoft JhengHei UI Light", 0, 16)); // NOI18N
        lblInput.setText("Input Directory:");
        lblInput.setOpaque(true);

        txtInputDir.setFont(new java.awt.Font("Segoe UI Light", 0, 16)); // NOI18N

        lblOutput.setFont(new java.awt.Font("Microsoft JhengHei UI Light", 0, 16)); // NOI18N
        lblOutput.setText("Output Directory:");
        lblOutput.setOpaque(true);

        txtOutputDir.setFont(new java.awt.Font("Segoe UI Light", 0, 16)); // NOI18N

        btnInput.setText("...");
        btnInput.setToolTipText("Opens a dialog to select input file");
        btnInput.setFocusable(false);
        btnInput.addMouseListener(new java.awt.event.MouseAdapter() {
            public void mouseClicked(java.awt.event.MouseEvent evt) {
                btnInputMouseClicked(evt);
            }
        });

        btnOutput.setText("...");
        btnOutput.setToolTipText("Opens a dialog to select output directory");
        btnOutput.setFocusable(false);
        btnOutput.addMouseListener(new java.awt.event.MouseAdapter() {
            public void mouseClicked(java.awt.event.MouseEvent evt) {
                btnOutputMouseClicked(evt);
            }
        });

        lblASL.setFont(new java.awt.Font("Microsoft JhengHei UI Light", 0, 16)); // NOI18N
        lblASL.setText("ASL Compiler Directory");
        lblASL.setOpaque(true);

        txtASLDir.setFont(new java.awt.Font("Segoe UI Light", 0, 16)); // NOI18N

        btnASL.setText("...");
        btnASL.setToolTipText("Opens a dialog to select the compiler location");
        btnASL.setFocusable(false);
        btnASL.addMouseListener(new java.awt.event.MouseAdapter() {
            public void mouseClicked(java.awt.event.MouseEvent evt) {
                btnASLMouseClicked(evt);
            }
        });

        jSeparator.setToolTipText("");

        lblASLSmall.setFont(new java.awt.Font("Microsoft YaHei UI", 0, 10)); // NOI18N
        lblASLSmall.setText("Location of the asl.exe file.");
        lblASLSmall.setOpaque(true);

        lblInputSmall.setFont(new java.awt.Font("Microsoft YaHei UI", 0, 10)); // NOI18N
        lblInputSmall.setText("Directory of scripts that will be compiled in to the output directory.");
        lblInputSmall.setOpaque(true);

        lblOutputSmall.setFont(new java.awt.Font("Microsoft YaHei UI", 0, 10)); // NOI18N
        lblOutputSmall.setText("Directory that the compiled .sqf script(s) will be saved in.");
        lblOutputSmall.setOpaque(true);

        cbCompileAll.setFont(new java.awt.Font("Microsoft YaHei UI", 0, 11)); // NOI18N
        cbCompileAll.setText("Compile all scripts including sub folders.");
        cbCompileAll.setFocusable(false);
        cbCompileAll.addChangeListener(new javax.swing.event.ChangeListener() {
            public void stateChanged(javax.swing.event.ChangeEvent evt) {
                cbCompileAllStateChanged(evt);
            }
        });

        cbPrettyPrinting.setFont(new java.awt.Font("Microsoft YaHei UI", 0, 11)); // NOI18N
        cbPrettyPrinting.setText("Activate pretty printing.");
        cbPrettyPrinting.setFocusable(false);
        cbPrettyPrinting.addChangeListener(new javax.swing.event.ChangeListener() {
            public void stateChanged(javax.swing.event.ChangeEvent evt) {
                cbPrettyPrintingStateChanged(evt);
            }
        });

        btnCompile.setFont(new java.awt.Font("Microsoft JhengHei UI Light", 0, 16)); // NOI18N
        btnCompile.setText("Compile");
        btnCompile.setToolTipText("Opens a dialog to select output directory");
        btnCompile.setFocusable(false);
        btnCompile.addMouseListener(new java.awt.event.MouseAdapter() {
            public void mouseClicked(java.awt.event.MouseEvent evt) {
                btnCompileMouseClicked(evt);
            }
        });

        lblASLError.setFont(new java.awt.Font("Microsoft YaHei UI", 0, 10)); // NOI18N
        lblASLError.setForeground(java.awt.Color.red);
        lblASLError.setText("Some error");
        lblASLError.addMouseListener(new java.awt.event.MouseAdapter() {
            public void mouseClicked(java.awt.event.MouseEvent evt) {
                lblASLErrorMouseClicked(evt);
            }
        });

        lblInputError.setFont(new java.awt.Font("Microsoft YaHei UI", 0, 10)); // NOI18N
        lblInputError.setForeground(java.awt.Color.red);
        lblInputError.setText("Some error");

        lblOutputError.setFont(new java.awt.Font("Microsoft YaHei UI", 0, 10)); // NOI18N
        lblOutputError.setForeground(java.awt.Color.red);
        lblOutputError.setText("Some error");
        lblOutputError.addMouseListener(new java.awt.event.MouseAdapter() {
            public void mouseClicked(java.awt.event.MouseEvent evt) {
                lblOutputErrorMouseClicked(evt);
            }
        });

        javax.swing.GroupLayout layout = new javax.swing.GroupLayout(getContentPane());
        getContentPane().setLayout(layout);
        layout.setHorizontalGroup(
            layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addComponent(jSeparator)
            .addGroup(layout.createSequentialGroup()
                .addGap(15, 15, 15)
                .addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
                    .addGroup(layout.createSequentialGroup()
                        .addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
                            .addComponent(lblOutputError, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
                            .addComponent(lblInputError, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
                            .addComponent(lblASLError, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, Short.MAX_VALUE)
                            .addGroup(layout.createSequentialGroup()
                                .addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
                                    .addComponent(lblASL)
                                    .addComponent(lblASLSmall)
                                    .addGroup(layout.createSequentialGroup()
                                        .addComponent(txtASLDir, javax.swing.GroupLayout.PREFERRED_SIZE, 320, javax.swing.GroupLayout.PREFERRED_SIZE)
                                        .addGap(6, 6, 6)
                                        .addComponent(btnASL))
                                    .addGroup(layout.createSequentialGroup()
                                        .addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.TRAILING)
                                            .addComponent(txtOutputDir, javax.swing.GroupLayout.Alignment.LEADING, javax.swing.GroupLayout.PREFERRED_SIZE, 320, javax.swing.GroupLayout.PREFERRED_SIZE)
                                            .addComponent(lblOutput, javax.swing.GroupLayout.Alignment.LEADING))
                                        .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                                        .addComponent(btnOutput))
                                    .addComponent(lblOutputSmall)
                                    .addComponent(lblInput)
                                    .addGroup(layout.createSequentialGroup()
                                        .addComponent(txtInputDir, javax.swing.GroupLayout.PREFERRED_SIZE, 320, javax.swing.GroupLayout.PREFERRED_SIZE)
                                        .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                                        .addComponent(btnInput))
                                    .addComponent(lblInputSmall)
                                    .addComponent(cbCompileAll)
                                    .addComponent(cbPrettyPrinting))
                                .addGap(0, 0, Short.MAX_VALUE)))
                        .addGap(24, 24, 24))
                    .addGroup(layout.createSequentialGroup()
                        .addComponent(btnCompile, javax.swing.GroupLayout.PREFERRED_SIZE, 373, javax.swing.GroupLayout.PREFERRED_SIZE)
                        .addGap(22, 22, 22))))
        );
        layout.setVerticalGroup(
            layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
            .addGroup(layout.createSequentialGroup()
                .addGap(6, 6, 6)
                .addComponent(lblASL)
                .addGap(3, 3, 3)
                .addComponent(lblASLSmall)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
                    .addComponent(txtASLDir, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(btnASL, javax.swing.GroupLayout.PREFERRED_SIZE, 28, javax.swing.GroupLayout.PREFERRED_SIZE))
                .addGap(4, 4, 4)
                .addComponent(lblASLError)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addComponent(jSeparator, javax.swing.GroupLayout.PREFERRED_SIZE, 10, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addGap(3, 3, 3)
                .addComponent(lblInput)
                .addGap(3, 3, 3)
                .addComponent(lblInputSmall, javax.swing.GroupLayout.PREFERRED_SIZE, 14, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.RELATED)
                .addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
                    .addComponent(txtInputDir, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(btnInput, javax.swing.GroupLayout.PREFERRED_SIZE, 28, javax.swing.GroupLayout.PREFERRED_SIZE))
                .addGap(4, 4, 4)
                .addComponent(lblInputError)
                .addGap(6, 6, 6)
                .addComponent(lblOutput)
                .addGap(3, 3, 3)
                .addComponent(lblOutputSmall, javax.swing.GroupLayout.PREFERRED_SIZE, 14, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addGap(6, 6, 6)
                .addGroup(layout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
                    .addComponent(txtOutputDir, javax.swing.GroupLayout.PREFERRED_SIZE, javax.swing.GroupLayout.DEFAULT_SIZE, javax.swing.GroupLayout.PREFERRED_SIZE)
                    .addComponent(btnOutput, javax.swing.GroupLayout.PREFERRED_SIZE, 28, javax.swing.GroupLayout.PREFERRED_SIZE))
                .addGap(4, 4, 4)
                .addComponent(lblOutputError)
                .addGap(6, 6, 6)
                .addComponent(cbCompileAll)
                .addPreferredGap(javax.swing.LayoutStyle.ComponentPlacement.UNRELATED)
                .addComponent(cbPrettyPrinting)
                .addGap(11, 11, 11)
                .addComponent(btnCompile, javax.swing.GroupLayout.PREFERRED_SIZE, 41, javax.swing.GroupLayout.PREFERRED_SIZE)
                .addGap(11, 11, 11))
        );

        pack();
        setLocationRelativeTo(null);
    }// </editor-fold>//GEN-END:initComponents

    private void btnInputMouseClicked(java.awt.event.MouseEvent evt) {//GEN-FIRST:event_btnInputMouseClicked
        if (SwingUtilities.isLeftMouseButton(evt)) {
            String path = fileChooser("Select input directory", null);
            File inputDir = new File(path);
            if (inputDir.exists()) {
                prefs.put("inputDir", path);
                txtInputDir.setText(path);
            }
        }
    }//GEN-LAST:event_btnInputMouseClicked

    private void btnOutputMouseClicked(java.awt.event.MouseEvent evt) {//GEN-FIRST:event_btnOutputMouseClicked
        if (SwingUtilities.isLeftMouseButton(evt)) {
            String path = fileChooser("Select output directory", null);
            File outputDir = new File(path);
            if (outputDir.exists() && outputDir.isDirectory()) {
                prefs.put("outputDir", path);
                txtOutputDir.setText(path);
            } else if (!outputDir.exists()) {
                lblOutputError.setText("Output folder doesn't exsist! Click this message to create it.");
                outputDirFix = true;
            }
        }
    }//GEN-LAST:event_btnOutputMouseClicked

    private void btnASLMouseClicked(java.awt.event.MouseEvent evt) {//GEN-FIRST:event_btnASLMouseClicked
        if (SwingUtilities.isLeftMouseButton(evt)) {
            String path = fileChooser("Select 'asl.exe' location", new FileNameExtensionFilter("Executable", "exe"));
            File asl = new File(path);
            if (asl.exists()) {
                prefs.put("aslDir", path);
                txtASLDir.setText(path);
            }
        }
    }//GEN-LAST:event_btnASLMouseClicked

    private void btnCompileMouseClicked(java.awt.event.MouseEvent evt) {//GEN-FIRST:event_btnCompileMouseClicked
        if (SwingUtilities.isLeftMouseButton(evt)) {
            String sumthin = new File(txtASLDir.getText()).getParent() + "\\asl.exe";
            if (!new File(sumthin).exists()) {
                lblASLError.setText("asl.exe isn't in this location! You can click this message to download it.");
            }
            JTextField[] dirFields = {txtASLDir, txtInputDir, txtOutputDir};
            for (byte i = 0; i < 3; i++) {
                File bleh = new File(dirFields[i].getText());
                if (!bleh.exists()) {
                    gaveError = true;
                    switch (i) {
                        case 0:
                            lblASLError.setText("asl.exe isn't in this location! You can click this message to download it.");
                            aslFix = true;
                            break;
                        case 1:
                            lblInputError.setText("This folder doesn't exist!");
                            break;
                        case 2:
                            lblOutputError.setText("Output folder doesn't exsist! Click this message to create it.");
                            outputDirFix = true;
                            break;
                    }
                    return;
                }
                if (bleh.exists() && gaveError) {
                    switch (i) {
                        case 0:
                            lblASLError.setText(" ");
                            break;
                        case 1:
                            lblInputError.setText(" ");
                            break;
                        case 2:
                            lblOutputError.setText(" ");
                            outputDirFix = true;
                            break;
                    }
                }
            }
            if (gaveError) {
                return;
            }
            String compileAll = "", prettyPrinting = "";
            if (cbCompileAll.isSelected()) {
                compileAll = "-r";
            }
            if (cbPrettyPrinting.isSelected()) {
                prettyPrinting = "-pretty";
            }
            String asl = txtASLDir.getText();
            String input = txtInputDir.getText();
            String output = txtOutputDir.getText();
            String aslError = " ";
            try {
                Process aslProcess = new ProcessBuilder(asl, compileAll, prettyPrinting, input, output).start();
                InputStream is = aslProcess.getInputStream();
                InputStreamReader isr = new InputStreamReader(is);
                BufferedReader br = new BufferedReader(isr);
                String line;
                while ((line = br.readLine()) != null) {
                    if (line.contains("Error")) {
                        aslError = line;
                    }
                }
                aslProcess.waitFor();
                if (!aslError.equals(" ")) {
                    DlgError error = new DlgError(this, true, aslError);
                    error.setLocationRelativeTo(this);
                    error.setVisible(true);
                }
                //Runtime.getRuntime().exec("cmd /c start \"" + asl + "\" " + compileAll + prettyPrinting + "\"" + input + "\" \"" + output + "\"");
                //TODO: Start app in command line with parameters
            } catch (IOException | InterruptedException ex) {
                Logger.getLogger(Main.class.getName()).log(Level.SEVERE, null, ex);
            }
        }
    }//GEN-LAST:event_btnCompileMouseClicked

    private void cbCompileAllStateChanged(javax.swing.event.ChangeEvent evt) {//GEN-FIRST:event_cbCompileAllStateChanged
        prefs.putBoolean("compileAll", cbCompileAll.isSelected());
    }//GEN-LAST:event_cbCompileAllStateChanged

    private void cbPrettyPrintingStateChanged(javax.swing.event.ChangeEvent evt) {//GEN-FIRST:event_cbPrettyPrintingStateChanged
        prefs.putBoolean("prettyPrinting", cbPrettyPrinting.isSelected());
    }//GEN-LAST:event_cbPrettyPrintingStateChanged

    private void lblOutputErrorMouseClicked(java.awt.event.MouseEvent evt) {//GEN-FIRST:event_lblOutputErrorMouseClicked
        if (SwingUtilities.isLeftMouseButton(evt) && outputDirFix) {
            new File(txtOutputDir.getText()).mkdirs();
            outputDirFix = false;
            lblOutputError.setText(" ");
        }
    }//GEN-LAST:event_lblOutputErrorMouseClicked

    private void lblASLErrorMouseClicked(java.awt.event.MouseEvent evt) {//GEN-FIRST:event_lblASLErrorMouseClicked
        if (SwingUtilities.isLeftMouseButton(evt) && aslFix) {
            try {
                URI github = new URI("https://github.com/DeKugelschieber/asl/releases");
                java.awt.Desktop.getDesktop().browse(github);
            } catch (URISyntaxException | IOException ex) {
                Logger.getLogger(Main.class.getName()).log(Level.SEVERE, null, ex);
            }
            aslFix = false;
            lblASLError.setText(" ");
        }
    }//GEN-LAST:event_lblASLErrorMouseClicked

    public static void main(String args[]) {
        try {
            UIManager.setLookAndFeel("com.sun.java.swing.plaf.windows.WindowsLookAndFeel");
        } catch (ClassNotFoundException | InstantiationException | IllegalAccessException | UnsupportedLookAndFeelException ex) {
            Logger.getLogger(Main.class.getName()).log(Level.SEVERE, null, ex);
        }
        UIManager.put("ToolTip.background", new ColorUIResource(255, 255, 255));
        UIManager.put("ToolTip.foreground", new ColorUIResource(87, 87, 87));
        Border lineBorder = BorderFactory.createLineBorder(new Color(118, 118, 118));
        UIManager.put("ToolTip.border", lineBorder);
        Border compoundBorder = BorderFactory.createCompoundBorder(UIManager.getBorder("ToolTip.border"), BorderFactory.createEmptyBorder(0, 2, 2, 3));
        UIManager.put("ToolTip.border", compoundBorder);
        java.awt.EventQueue.invokeLater(() -> {
            new Main().setVisible(true);
        });
    }

    // Variables declaration - do not modify//GEN-BEGIN:variables
    private javax.swing.JButton btnASL;
    private javax.swing.JButton btnCompile;
    private javax.swing.JButton btnInput;
    private javax.swing.JButton btnOutput;
    private javax.swing.JCheckBox cbCompileAll;
    private javax.swing.JCheckBox cbPrettyPrinting;
    private javax.swing.JSeparator jSeparator;
    private javax.swing.JLabel lblASL;
    private javax.swing.JLabel lblASLError;
    private javax.swing.JLabel lblASLSmall;
    private javax.swing.JLabel lblInput;
    private javax.swing.JLabel lblInputError;
    private javax.swing.JLabel lblInputSmall;
    private javax.swing.JLabel lblOutput;
    private javax.swing.JLabel lblOutputError;
    private javax.swing.JLabel lblOutputSmall;
    private javax.swing.JTextField txtASLDir;
    private javax.swing.JTextField txtInputDir;
    private javax.swing.JTextField txtOutputDir;
    // End of variables declaration//GEN-END:variables
}
